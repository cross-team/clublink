import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { ShortLink } from '../../../entity/ShortLink';
import { UIFactory } from '../../UIFactory';
import { validateLongLinkFormat } from '../../../validators/LongLink.validator';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
import { raiseCreateShortLinkError } from '../../../state/actions';
import { IAppState } from '../../../state/reducers';
import { Store } from 'redux';
import { ShortLinkService } from '../../../service/ShortLink.service';
import { QrCodeService } from '../../../service/QrCode.service';

interface IProps {
  store: Store<IAppState>;
  uiFactory: UIFactory;
  shortLinkService: ShortLinkService;
  qrCodeService: QrCodeService;
  onShortLinkCreated?: (shortLink: string) => void;
  onAuthenticationFailed?: () => void;
}

interface IState {
  longLink: string;
  username: string;
  room: string;
  alias: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
  description: string;
  valid: null | boolean;
}

export class CreateShortLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      alias: '',
      longLink: '',
      username: '',
      room: '',
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: '',
      description: 'Enter the super-secret code 🤓',
      valid: true
    };
  }

  render(): ReactElement {
    return (
      <Section title={''}>
        <div className={'control create-short-link'}>
          <h1>
            <span className={'green'}>club</span>l
            <span className={'green'}>.</span>ink{' '}
            <span className={'slash'}>/</span>
          </h1>
          <div className={'text-field-wrapper'}>
            <TextField
              className="code"
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'enter code'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
              onFocus={this.handleFocus}
            />
          </div>
          {this.state.alias && (
            <span
              role="button"
              className={'rocket-button'}
              onClick={this.handleReserveShortLinkClick}
            >
              🚀
            </span>
          )}
        </div>
        <div className={'input-description'}>{this.state.description}</div>
        {this.state.valid === true && (
          <>
            <div className={'text-field-wrapper'}>
              <TextField
                text={this.state.longLink}
                placeHolder={
                  'Now enter your fricking ridiculously long shitty link here...'
                }
                onBlur={this.handleLongLinkTextFieldBlur}
                onChange={this.handleLongLinkChange}
              />
            </div>
            <div className={'text-field-wrapper'}>
              <TextField
                className="username"
                text={this.state.username}
                placeHolder={'@username'}
                onBlur={this.handleLongLinkTextFieldBlur}
                onChange={this.handleLongLinkChange}
              />
              <TextField
                className="room"
                text={this.state.room}
                placeHolder={'your room title'}
                onBlur={this.handleLongLinkTextFieldBlur}
                onChange={this.handleLongLinkChange}
              />
            </div>
            <div className={'input-error'}>{this.state.inputError}</div>
          </>
        )}
        {/* <div className="create-short-link-btn">
            <Button onClick={this.handleCreateShortLinkClick}>
              Create Short Link
            </Button>
          </div> */}
        {/* {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })} */}
        {this.state.shouldShowUsage && (
          <div className={'short-link-usage-wrapper'}>
            <ShortLinkUsage
              shortLink={this.state.createdShortLink}
              longLink={this.state.createdLongLink}
              qrCodeUrl={this.state.qrCodeURL}
            />
          </div>
        )}
      </Section>
    );
  }

  autoFillInLongLink(longLink: string) {
    if (!longLink) {
      return;
    }

    this.setState({
      longLink: longLink
    });

    const inputError = validateLongLinkFormat(longLink);
    if (inputError != null) {
      this.setState({
        inputError: inputError
      });
      return;
    }

    this.focusShortLinkTextField();
  }

  handleFocus = () => {
    this.setState({
      description: 'Keep it simple, it only lasts 24 hours ✌️'
    });
  };

  handleLongLinkTextFieldBlur = () => {
    const { longLink } = this.state;
    const err = validateLongLinkFormat(longLink);
    this.setState({
      inputError: err || undefined
    });
  };

  handleLongLinkChange = (newLongLink: string) => {
    this.setState({
      longLink: newLongLink
    });
  };

  handleAliasChange = (newAlias: string) => {
    this.setState({
      alias: newAlias
    });
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    this.setState({
      inputError: err || undefined,
      description: 'Enter the super-secret code 🤓'
    });
  };

  handleReserveShortLinkClick = () => {
    const { alias } = this.state;
    const shortLink: ShortLink = {
      longLink: '#',
      alias: alias || ''
    };
    this.props.shortLinkService
      .createShortLink(shortLink, this.state.isShortLinkPublic)
      .then(async (createdShortLink: ShortLink) => {
        const shortLink = this.props.shortLinkService.aliasToFrontendLink(
          createdShortLink.alias!
        );

        const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

        this.setState({
          createdShortLink: shortLink,
          qrCodeURL: qrCodeURL,
          shouldShowUsage: true
        });

        if (this.props.onShortLinkCreated) {
          this.props.onShortLinkCreated(shortLink);
        }
      })
      .catch(({ authenticationErr, createShortLinkErr }) => {
        if (authenticationErr) {
          if (this.props.onAuthenticationFailed) {
            this.props.onAuthenticationFailed();
          }
          return;
        }
        this.props.store.dispatch(
          raiseCreateShortLinkError(createShortLinkErr)
        );
      });
  };

  handleCreateShortLinkClick = () => {
    const { alias, longLink } = this.state;
    const shortLink: ShortLink = {
      longLink: longLink,
      alias: alias || ''
    };
    this.props.shortLinkService
      .createShortLink(shortLink, this.state.isShortLinkPublic)
      .then(async (createdShortLink: ShortLink) => {
        const shortLink = this.props.shortLinkService.aliasToFrontendLink(
          createdShortLink.alias!
        );

        const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);

        this.setState({
          createdShortLink: shortLink,
          createdLongLink: longLink,
          qrCodeURL: qrCodeURL,
          shouldShowUsage: true
        });

        if (this.props.onShortLinkCreated) {
          this.props.onShortLinkCreated(shortLink);
        }
      })
      .catch(({ authenticationErr, createShortLinkErr }) => {
        if (authenticationErr) {
          if (this.props.onAuthenticationFailed) {
            this.props.onAuthenticationFailed();
          }
          return;
        }
        this.props.store.dispatch(
          raiseCreateShortLinkError(createShortLinkErr)
        );
      });
  };

  handlePublicToggleClick = (enabled: boolean) => {
    this.setState({
      isShortLinkPublic: enabled
    });
  };

  focusShortLinkTextField = () => {
    if (!this.shortLinkTextField.current) {
      return;
    }
    this.shortLinkTextField.current.focus();
  };
}

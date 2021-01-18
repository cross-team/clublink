import React, { Component, ReactElement } from 'react';

import './CreateShortLinkSection.scss';
import { TextField } from '../../form/TextField';
import { Button } from '../../ui/Button';
import { ShortLinkUsage } from './ShortLinkUsage';
import { Section } from '../../ui/Section';
import { ShortLink } from '../../../entity/ShortLink';
import { UIFactory } from '../../UIFactory';
import { validateCustomAliasFormat } from '../../../validators/CustomAlias.validator';
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
  alias?: string;
  inputError?: string;
  isShortLinkPublic?: boolean;
  shouldShowUsage: boolean;
  createdShortLink: string;
  createdLongLink: string;
  qrCodeURL: string;
}

export class VisitLinkSection extends Component<IProps, IState> {
  private shortLinkTextField = React.createRef<TextField>();

  constructor(props: IProps) {
    super(props);
    this.state = {
      shouldShowUsage: false,
      createdShortLink: '',
      createdLongLink: '',
      qrCodeURL: ''
    };
  }

  render(): ReactElement {
    return (
      <Section title={''}>
        <div className={'control visit-short-link'}>
          <h1>
            <span className={'green'}>club</span>l
            <span className={'green'}>.</span>ink{' '}
            <span className={'slash'}>/</span>
          </h1>
          <div className={'text-field-wrapper'}>
            <TextField
              ref={this.shortLinkTextField}
              text={this.state.alias}
              placeHolder={'enter code'}
              onBlur={this.handleCustomAliasTextFieldBlur}
              onChange={this.handleAliasChange}
            />
          </div>
        </div>
        <div className={'input-error'}>{this.state.inputError}</div>
        <div className={'input-description'}>
          Enter the super-secret code and go 🚀
        </div>
        {this.props.uiFactory.createPreferenceTogglesSubSection({
          uiFactory: this.props.uiFactory,
          isShortLinkPublic: this.state.isShortLinkPublic,
          onPublicToggleClick: this.handlePublicToggleClick
        })}
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

  handleAliasChange = (newAlias: string) => {
    this.setState({
      alias: newAlias
    });
  };

  handleCustomAliasTextFieldBlur = () => {
    const { alias } = this.state;
    const err = validateCustomAliasFormat(alias);
    this.setState({
      inputError: err || undefined
    });
  };

  handleCreateShortLinkClick = () => {
    // const { alias, longLink } = this.state;
    // const shortLink: ShortLink = {
    //   longLink: longLink,
    //   alias: alias || ''
    // };
    // this.props.shortLinkService
    //   .createShortLink(shortLink, this.state.isShortLinkPublic)
    //   .then(async (createdShortLink: ShortLink) => {
    //     const shortLink = this.props.shortLinkService.aliasToFrontendLink(
    //       createdShortLink.alias!
    //     );
    //     const qrCodeURL = await this.props.qrCodeService.newQrCode(shortLink);
    //     this.setState({
    //       createdShortLink: shortLink,
    //       createdLongLink: longLink,
    //       qrCodeURL: qrCodeURL,
    //       shouldShowUsage: true
    //     });
    //     if (this.props.onShortLinkCreated) {
    //       this.props.onShortLinkCreated(shortLink);
    //     }
    //   })
    //   .catch(({ authenticationErr, createShortLinkErr }) => {
    //     if (authenticationErr) {
    //       if (this.props.onAuthenticationFailed) {
    //         this.props.onAuthenticationFailed();
    //       }
    //       return;
    //     }
    //     this.props.store.dispatch(
    //       raiseCreateShortLinkError(createShortLinkErr)
    //     );
    //   });
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
